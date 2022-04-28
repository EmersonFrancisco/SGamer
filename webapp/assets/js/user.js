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
        window.location = userId;
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
        window.location = userId;
    }).fail(function() {
        Swal.fire('Opps....', 'erro ao seguir usuário!', 'error');
    }).always(function() {
        element.prop('disabled', false);
    });

}