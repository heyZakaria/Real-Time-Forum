
export const myCode = {
    home: `
    <!--                 HOME                   -->
        <div class="homeCode">
            <div>
                <form method="post" class="post_form">
                    <label for="title">Title</label>
                    <input class="post_title" name="title" type="text" placeholder="Name it" maxlength="100">
                    <label for="content">Content</label>
                    <input class="post_content" name="content" type="text" placeholder="Speak your mind soldier"
                        maxlength="1000">
                    <fieldset>
                        <legend>Select an appropriate category</legend>
                        <label for="sport">Sport</label>
                        <input class="category" type="checkbox" value="sport" id="sport" name="category" />
                        <label for="science">Science</label>
                        <input class="category" type="checkbox" id="science" name="category" value="science">
                        <label for="entertainment">Entertainment</label>
                        <input class="category" type="checkbox" id="entertainment" name="category"
                            value="entertainment">
                    </fieldset>
                    <button class="post_btn" type="button">Post</button>
                </form>
                <div class="filter_form">
                    <form action="">
                        <select class="selectfilter" name="filter" multiple>
                            <option value="sport">Sport</option>
                            <option value="science">Science</option>
                            <option value="entertainment">Entertainment</option>
                            <option value="liked">Liked</option>
                            <option value="created">Created</option>
                        </select>
                        <input class="filterbutton" type="button" value="Filter">
                    </form>
                </div>
                <div class="container">
                    <div id="posts">
                    </div>
                    <div class="pagination">
                    </div>
                </div>
            </div>
        </div>
`,

    login: `
    <!--                 LOGIN                   -->
        <div class="loginCode">
            <div class="form-container">
                <h2>Login</h2>
                <form method="post" class="login_form">
                    <label for="email">EMAIL OR USERNAME:</label>
                    <input type="text" id="login_email" name="emailorusername" placeholder="Email or Username" required>
                    <label for="login_password">PASSWORD:</label>
                    <input type="password" id="login_password" name="login_password" placeholder="Setup your password"
                        required>
                    <label></label>
                    <input type="submit" value="Login">
                    <span id="server_error"></span>
                </form>
            </div>
        </div>`,

    register: `
    <!--                 REGISTER                   -->
        <div class="registerCode">
            <div class="form-container">
                <h2>Register</h2>
                <form class="regsiter_form" method="post">
                    <label for="username">USER NAME</label>
                    <input type="text" id="username" name="username" placeholder="Username">
                    <br>
                    <span id="error0"></span>
                    <label for="register_email">EMAIL *</label>
                    <input type="email" id="register_email" name="register_email" placeholder="Enter your email here">
                    <br>
                    <span id="error1"></span>
                    <label for="register_password">PASSWORD *</label>
                    <input type="password" id="register_password" name="register_password"
                        placeholder="Setup your password">
                    <br>
                    <span id="error2"></span>
                    <label for="register_password_2">REPEAT YOUR PASSWORD *</label>
                    <input type="password" id="register_password_2" name="register_password_2"
                        placeholder="Repeat your password">
                    <br>
                    <span id="error3"></span>
                    <input type="submit" id="register_button" value="Register">
                    <div id="server_error"></div>
                </form>
            </div>
        </div>`,

    errata: `<!-- ERRROOOOOOOOOOOOR -->
        <div class="errorCode">

            <main class="err-main">
                <div class="err-bg"></div>
                <div class="err-content">
                    <h1>ERROOOOOOR</h1>
                    <!-- <form action="/"> -->
                    <button class="go-back-error" onclick="window.history.back()">Go Back</button>
                    <!-- </form> -->
                </div>
            </main>
        </div>`,
}