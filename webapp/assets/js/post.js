$('#newPost').on('submit', newPost);
$('#updatePost').on('click', updatePost);
$('.deletePost').on('click', deletePost);
$(document).on('click', '.likePost', likePost);
$(document).on('click', '.unlikePost', unlikePost);


function newPost(evento) {
    evento.preventDefault();
    $(this).prop('disable', true);
    $.ajax({
            url: "/post",
            method: "POST",
            data: {
                title: $('#title').val(),
                content: $('#content').val(),
            }
        })
        .done(function() {
            window.location = "/home";
        }).fail(function() {
            Swal.fire('Opps....', 'Erro ao Salvar Publicação', 'error');
        });
}

function likePost(evento) {
    evento.preventDefault();
    const element = $(evento.target);
    const clickerElement = evento.target || evento.srcElement;
    const postId = clickerElement.id;

    element.prop('disabled', true)
    $.ajax({
        url: `/post/${postId}/like`,
        method: "POST",

    }).done(function() {
        const countLike = element.next('span');
        const quantLike = parseInt(countLike.text());

        countLike.text(quantLike + 1);

        element.addClass('unlikePost');
        element.addClass('text-danger');
        element.removeClass('likePost');
    }).fail(function() {
        Swal.fire('Opps....', 'erro ao curtir publicação!', 'error');
    }).always(function() {
        element.prop('disabled', false);
    });

}

function unlikePost(evento) {
    evento.preventDefault();
    const element = $(evento.target);
    const clickerElement = evento.target || evento.srcElement;
    const postId = clickerElement.id;

    element.prop('disabled', true)
    $.ajax({
        url: `/post/${postId}/unlike`,
        method: "POST",
    }).done(function() {
        const countLike = element.next('span');
        const quantLike = parseInt(countLike.text());

        countLike.text(quantLike - 1);

        element.addClass('likePost');
        element.removeClass('text-danger');
        element.removeClass('unlikePost');
    }).fail(function() {
        Swal.fire('Opps....', 'Erro ao descurtir Publicação', 'error');
    }).always(function() {
        element.prop('disabled', false);
    });

}

function updatePost(evento) {
    evento.preventDefault();

    $(this).prop('disable', true);
    const postId = $(this).attr('data-postId');
    $.ajax({
        url: `/post/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function() {
        Swal.fire('Sucesso!', 'Publicação Criada com sucesso!', 'success').then(function() {
            window.location = "/home";
        });

    }).fail(function() {
        Swal.fire('Opps....', 'Erro ao atualizar a publicação!', 'error');
    }).always(function() {
        $('#updatePost').prop('disabled', false);
    });
}

function deletePost(evento) {
    evento.preventDefault();

    Swal.fire({
        title: 'Deseja deletar?',
        text: "você realmente quer deletar essa publicação?",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Deletar!',
        cancelButtonText: 'Cancelar'
    }).then((result) => {
        if (result.isConfirmed) {
            Swal.fire(
                'Deletado!',
                'Sua publicação foi deletada.',
                'success'
            ).then(function() {
                const element = $(evento.target);
                const clickerElement = evento.target || evento.srcElement;
                const post = element.closest('div');
                const postId = clickerElement.id;
                element.prop('disabled', true)
                $.ajax({
                    url: `/post/${postId}`,
                    method: "DELETE",
                }).done(function() {
                    post.fadeOut("slow", function() {
                        $(this).remove();
                    });
                }).fail(function() {
                    Swal.fire('Opps....', 'Erro ao deletar a publicação!', 'error');
                }).always(function() {
                    element.prop('disabled', false);
                });
            })
        }
    })
}