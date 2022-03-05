<script lang="ts">
    import { Comment, Favorite, SurfSession } from "$lib/utils/models"
    import { favorites, addFavorite, removeFavorite } from "$lib/stores/favoritesStore";
    import { comments, getComments, postComment } from "$lib/stores/commentsStore";
    import { user, users, isAuthenticated } from "$lib/stores/userStore"; 
    import { slide } from "svelte/transition";
    import { onMount } from "svelte";
    import IconCommentOutline from "~icons/mdi/comment-outline"
    import IconCardsHeartOutline from "~icons/mdi/cards-heart-outline"
    import IconCardsHeart from "~icons/mdi/cards-heart"

    export let session: SurfSession;

    let commentsOnPost: Comment[] = [];
    let favoritesOnPost: Favorite[] = [];

    let newComment: string;

    let postIsLiked = false;

    let displayComments = false;

    onMount(async () => {
        if(await isAuthenticated()){
            postIsLiked = $favorites.filter(fav => fav.surf_session_id == session.id && fav.user_id == $user.id).length > 0
        }
    })

    let displayLoginToLikeMessage = false;
    let displayLoginToCommentMessage = false;

      
    $: if($comments.length > 0){
            commentsOnPost = $comments.filter(comment => comment.surf_session_id == session.id)
        }
        
    $: if($favorites.length >= 0){
            favoritesOnPost = $favorites.filter(favorite => favorite.surf_session_id == session.id)
        }

</script>

<main>
    <div class="surf-session-container">
        <div class="header-container">
            <div class="session-spot">{session.surf_spot}</div>
            <div class="date">{session.datetime}</div>
        </div>
        <div class="session-location">{session.spot_location}</div>
        
        
        <div class="session-conditions">
            <div class="swell">Swell: {session.swell}</div>
            <div class="wind">Wind: {session.wind}</div>
            <div class="tide">Tide: {session.tide}</div>
        </div>
        <div class="session-description">{session.description}</div>
        {#if session.surfboard_id}
            <div class="session-surfboard"></div>
        {/if}
        <div class="bottom-line">
            <div class="favorite-comment-buttons">
                <div class="favorites" on:click={async () => {
                    if(await isAuthenticated()){
                        if(!postIsLiked){
                            addFavorite(session.id, $user.id)
                            postIsLiked = true
                        } else {
                            removeFavorite($favorites.filter(fav => fav.surf_session_id == session.id && fav.user_id == $user.id)[0].id)
                            postIsLiked = false
                        } 
                    } else {
                        displayLoginToLikeMessage = true;
                    }
                }}>
                    {#if !postIsLiked}
                        <div><IconCardsHeartOutline style="font-size: 1.4em"/></div>
                    {/if}
                    {#if postIsLiked}
                        <div><IconCardsHeart style="font-size: 1.4em"/></div>
                    {/if}
                    <div class="favorite-count">{favoritesOnPost.length}</div>
                </div>
                <div class="comments" on:click={() => displayComments = !displayComments} >
                    <div><IconCommentOutline style="font-size: 1.35em; margin-top: 1px"/></div>
                    <div class="comment-count">{commentsOnPost.length}</div>
                </div>
                
            </div>
            <div class="user">{$users.filter(user => user.id == session.creator_id)[0].username}</div>
        </div>
        {#if displayLoginToLikeMessage}
            <div class="must-login">Login to favorite sessions</div>
        {/if}
        {#if displayComments}
            <div class="comment-section-container" transition:slide="{{ duration: 400 }}">
                {#if commentsOnPost.length > 0}
                    <div class="comments-container">
                        {#each commentsOnPost as comment}
                            <div class="comment">{$users.filter(user => user.id = comment.commenter_id)[0].username}: {comment.content}</div>
                        {/each}
                    </div>
                {/if}
                <textarea type="text" class="comment-input" placeholder="Add a Comment..." bind:value={newComment}></textarea>
                <button class="post-comment-button" on:click={async () => {
                    if(await isAuthenticated()){
                        if(newComment.length > 0){
                            postComment({
                                surf_session_id: session.id,
                                content: newComment,
                                commenter_id: $user.id
                            })
                            newComment = ""
                        }
                    } else {
                        displayLoginToCommentMessage = true
                    }
                }}>
                    Post
                </button>
                {#if displayLoginToCommentMessage}
                    <div class="must-login">Login to post comments</div>
                {/if}
            </div>
        {/if}
       
        
    </div>
</main>

<style>
     *, *::before, *::after {
		box-sizing: border-box;
		margin: 0;
		padding: 0;
	}

     .surf-session-container {
        background-color: #313131;
        margin-bottom: 1em;
        padding: 0.5em 1em;
        color: white;
        width: 450px;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
    }

    .header-container {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 0.2em;
    }

    .session-spot {
        font-size: 2em;
        padding-right: 00.75em;
        font-family: Verdana, sans-serif;
    }

    .date {
        font-size: 1em;
        font-family: Verdana, sans-serif;
        text-align: center;
        width: 123px;
        align-self: flex-start;
    }

    .session-location {
        font-size: 1.4em;
        margin-bottom: 0.4em;
        font-family: Verdana, sans-serif;
    }

    .user {
        font-size: 1.2em;
        margin-bottom: 00.3em;
        font-family: Verdana, sans-serif;
    }

    .session-conditions {
        margin-bottom: 00.6em;
        font-size: 1.1em;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
    }

    .swell, .wind {
        margin-bottom: 0.2em;
    }

    .session-description {
        background-color: #111111;
        border-radius: 5px;
        padding: 0.4em 1em;
        margin-bottom: 0.8em;
        font-family: Verdana, sans-serif;
    }

    .bottom-line {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
    }

    .must-login {
        margin-top: 0.25em;
        font-family: Helvetica;
        font-weight: lighter;
    }

    .favorite-comment-buttons {
        display: flex;
        justify-content: space-between;
        width: 15%;
        margin-bottom: 0.25em;
    }



    .favorites, .comments {
        display: flex;
        flex-direction: column;
        align-items: center;
        cursor: pointer;
        font-family: Helvetica, sans-serif;
        font-weight: lighter;
    }

    .favorite-count, .comment-count {
        font-family: Arial, sans-serif;
        font-weight: lighter;
        color: white;
        font-size: 1.2em;
    }

    .comment-count {
        margin-top: -0.8px;
    }

    .comments-container {
        background-color: #111111;
        border-radius: 5px;
        padding: 0.5em 1em;
        margin: 0.5em 0;
    }

    .comment {
        font-family: Verdana, Geneva, Tahoma, sans-serif;
        font-size: 00.85em;
    }

    button {
        position: relative;
        background: linear-gradient(to right bottom, #6a37c2, #49329e);
        color: rgb(240, 234, 234);
        padding: 0.5em 1em;
        text-align: center;
        font-size: 1em;
        font-family: Verdana, sans-serif;
        border: none;
        border-radius: 5px;
        margin: 0em 0 0.5em 0;
        box-shadow: rgba(0, 0, 0, 0.3) 0px 4px 12px;
        cursor: pointer;
        transition-duration: 300ms;
        z-index: 1;
    }

    button::before{
		position: absolute;
		content: "";
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		border-radius: 5px;
		background: linear-gradient(to right bottom, rgb(59, 31, 110),rgb(57, 39, 122)0);
		transition: opacity 0.3s;
		z-index: -1;
		opacity: 0;
	}

    button:hover::before {
		opacity: 1;
	}

    button:hover {
        color: rgb(182, 182, 182);
    }

    textarea {
        font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
        font-size: 1em;
        padding: 0.5em 1em;
        resize: none;
        background-color: #111111;
        color: white;
        border: none;
        border-radius: 5px;
        width: 100%;
        margin: 0.5em 0 0.5em 0;
    }

    @media (max-width: 600px) {
        .surf-session-container {
            width: 85vw;
        }

        .favorite-comment-buttons {
            width: 60px;
        }

        .header-container {
            align-items: flex-start;
            margin-bottom: 0;
        }

        .session-spot {
            font-size: 1.5em;
        }

        .date {
            font-size: 0.9em;
        }

        .session-location {
            font-size: 1.2em;
        }

        .session-conditions {
            font-size: 0.95em;
        }

        .session-description {
            font-size: 00.95em;
        }
        .user {
            font-size: 1em;
        }
    }
</style>