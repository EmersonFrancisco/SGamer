$('#registration-form').on('submit', newUser);


function newUser(evento) {
    evento.preventDefault();
    if ($('#pass').val() != $('#passConfirm').val()) {
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/user",
        method: "POST",
        data: {
            username: $('#username').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            pass: $('#pass').val(),
        }
    }).done(function() {
        window.location = "/login"
        alert("Usuário cadastrado com Sucesso!!");
    }).fail(function(erro) {
        console.log(erro)
        alert("Ocorreu um erro ao efeutar o cadastro!!");
    });
}