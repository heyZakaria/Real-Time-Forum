import { Registred } from "./registred.js"

export async function reactComment(postsData) {


    let like_button = document.querySelector(`.comment_like_${postsData.id}`)


    let dislike_button = document.querySelector(`.comment_dislike_${postsData.id}`)

    like_button.onclick = async function () {

        let userid = await Registred()

        if (!userid) {
            window.location.replace("/login");
        } else {


            let bol = false

            if (postsData.likers != null) {
                for (let e of postsData.likers) {
                    if (e == userid) {
                        bol = true
                        break
                    }

                }
            }



            let isDisliked = false

            if (postsData.dislikers != null) {
                for (let e of postsData.dislikers) {
                    if (e == userid) {
                        isDisliked = true
                        break
                    }
                }
            }

            try {

                const action = !bol ? 'increment' : 'decrement';
                if ((action == 'increment') && (isDisliked == true)) {

                    const result1 = await updateLikeandDislikeComments(postsData.id, "Dislike", 'decrement');
                    if (result1.success) {
                        postsData.dislike -= 1;
                        const index = postsData.dislikers.indexOf(userid);

                        if (index > -1) {

                            postsData.dislikers.splice(index, 1);
                        }


                    }
                    dislike_button.innerHTML = `dislikes ${postsData.dislike}`;
                }


                const result = await updateLikeandDislikeComments(postsData.id, "like", action);

                if (result.success) {

                    if (!bol) {
                        postsData.like += 1;
                        //isLiked = true;

                        if (!Array.isArray(postsData.likers)) {
                            postsData.likers = [];
                        }
                        postsData.likers.push(userid);
                    } else {
                        postsData.like -= 1;
                        const index = postsData.likers.indexOf(userid);

                        if (index > -1) {

                            postsData.likers.splice(index, 1);
                        }

                    }

                    like_button.innerHTML = `Likes ${postsData.like}`;
                }
            } catch (error) {
                console.error('Failed   update  like or inlike:', error);
            }
        }
    }


    dislike_button.onclick = async function () {

        let userid = await Registred()

        if (!userid) {
            window.location.replace("/login");
        } else {



            let bol = false

            if (postsData.dislikers != null) {
                for (let e of postsData.dislikers) {
                    if (e == userid) {
                        bol = true
                        break
                    }

                }
            }
            let isliked = false

            if (postsData.likers != null) {
                for (let e of postsData.likers) {
                    if (e == userid) {
                        isliked = true
                        break
                    }
                }
            }



            try {
                const action = !bol ? 'increment' : 'decrement';
                if ((action == 'increment') && (isliked == true)) {
                    const result1 = await updateLikeandDislikeComments(postsData.id, "like", 'decrement');
                    if (result1.success) {
                        postsData.like -= 1;
                        const index = postsData.likers.indexOf(userid);

                        if (index > -1) {

                            postsData.likers.splice(index, 1);
                        }
                    }
                    like_button.innerHTML = `Likes ${postsData.like}`;
                }
                const result = await updateLikeandDislikeComments(postsData.id, "Dislike", action);

                if (result.success) {
                    if (!bol) {
                        postsData.dislike += 1;
                        if (!Array.isArray(postsData.dislikers)) {
                            postsData.dislikers = [];
                        }
                        //   isDisliked = true;
                        postsData.dislikers.push(userid)
                    } else {

                        postsData.dislike -= 1;
                        const index = postsData.dislikers.indexOf(userid);

                        if (index > -1) {

                            postsData.dislikers.splice(index, 1);
                        }


                    }
                    dislike_button.innerHTML = `dislikes ${postsData.dislike}`;
                }
            } catch (error) {
                console.error('Failed   update  dislikes or indislike:', error);
            }
        }

    }



}







async function updateLikeandDislikeComments(commentId, reaction, action) {

    if (reaction == "like") {

        try {
            const response = await fetch(`/api/comments/${commentId}/like`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    action: action // "crimenter ola dicrementer"
                })
            });

            if (!response.ok) {
                throw new Error('Failed  e like');
            }


            return await response.json();
        } catch (error) {

            console.error('Error updating like:', error);
            throw error;
        }

    } else {

        try {
            const response = await fetch(`/api/comments/${commentId}/dislike`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    action: action // "crimenter ola dicrementer"
                })
            });

            if (!response.ok) {
                throw new Error('Failed  e dislike');
            }

            return await response.json();
        } catch (error) {
            console.error('Error updating dislike:', error);
            throw error;
        }

    }


}