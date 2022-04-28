$('#unfollow').on('click', unFollow);
$('#follow').on('click', follow);

function unFollow(evento) {
    evento.preventDefault();
    const element = $(evento.target);
    userId = $(this).data("userid")
    element.prop('disabled', true);
    $.ajax({
        url: `/user/${userId}/unfollow`,
        method: "POST",
    }).done(function() {
        Swal.fire('Sucesso', 'deixou de seguir com sucesso!', 'success').then(function() {
            window.location = userId;
        })
    }).fail(function() {
        Swal.fire('Opps....', 'erro ao deixar de seguir usuário!', 'error');
    }).always(function() {
        element.prop('disabled', false);
    });

}

function follow(evento) {
    evento.preventDefault();
    const element = $(evento.target);
    userId = $(this).data("userid");
    element.prop('disabled', true);
    $.ajax({
        url: `/user/${userId}/follow`,
        method: "POST",

    }).done(function() {
        Swal.fire('Sucesso', 'Usuário seguido com sucesso!', 'success').then(function() {
            window.location = userId;
        })
    }).fail(function() {
        Swal.fire('Opps....', 'erro ao seguir usuário!', 'error');
    }).always(function() {
        element.prop('disabled', false);
    });

}