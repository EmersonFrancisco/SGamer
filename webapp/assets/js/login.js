$('#login').on('submit', login);

function login(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            pass: $('#pass').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Usuário ou senha inválida!");
    });
}