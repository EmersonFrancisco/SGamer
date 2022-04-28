$('#registration-form').on('submit', newUser);


function newUser(evento) {
    evento.preventDefault();
    if ($('#pass').val() != $('#passConfirm').val()) {
        Swal.fire('Opps....', 'As senhas não conhecidem!', 'error');
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
        Swal.fire('Sucesso!', 'Usuário cadastrado com sucesso!', 'success').then(function() {
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
                Swal.fire('Opps....', 'Erro ao autenticar o usuário!', 'error');
            });
        });
    }).fail(function(erro) {
        Swal.fire('Opps....', 'Erro ao cadastrar usuario!', 'error');
    });
}