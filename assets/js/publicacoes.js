$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descutir-publicacao', descurtirPublicacao);
$('#atualizar-publicacao').on('submit', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao(e) {
    e.preventDefault();
    console.log("Iniciando criação de publicação...");

    $.ajax({
        url: '/publicacoes',
        method: 'POST',
        contentType: 'application/json',
        data : {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(res => {
        window.location = "/home";
    }).fail(fail => {
        Swal.fire({
            title: 'Erro ao criar publicação',
            text: fail.responseText,
            icon: 'error',
            confirmButtonText: 'OK'
        });
    })
}

function curtirPublicacao(e) {
    e.preventDefault();
    const elementoClicado = $(e.currentTarget);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled', true);
    console.log("Curtindo publicação ID: " + publicacaoId);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: 'POST'
    }).done(res => {
        const contadorCurtidas = elementoClicado.next('span');
        let curtidasAtuais = parseInt(contadorCurtidas.text());
        contadorCurtidas.text(curtidasAtuais + 1);

        elementoClicado.addClass('descutir-publicacao').addClass('text-danger').removeClass('curtir-publicacao');
        elementoClicado.off('click').on('click', descurtirPublicacao);
    }).fail(fail => {
        Swal.fire({
            title: 'Erro ao curtir publicação',
            text: fail.responseText,
            icon: 'error',
            confirmButtonText: 'OK'
        });
    }).always(() => {
        elementoClicado.prop('disabled', false);
    });
}

function descurtirPublicacao(e) {
    e.preventDefault();
    const elementoClicado = $(e.currentTarget);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled', true);
    console.log("Descurtindo publicação ID: " + publicacaoId);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: 'POST'
    }).done(res => {
        const contadorCurtidas = elementoClicado.next('span');
        let curtidasAtuais = parseInt(contadorCurtidas.text());
        contadorCurtidas.text(curtidasAtuais - 1);
        elementoClicado.addClass('curtir-publicacao').removeClass('descutir-publicacao').removeClass('text-danger');
        elementoClicado.off('click').on('click', curtirPublicacao);
    }).fail(fail => {
        Swal.fire({
            title: 'Erro ao descurtir publicação',
            text: fail.responseText,
            icon: 'error',
            confirmButtonText: 'OK'
        });
    }).always(() => {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao(e) {
    e.preventDefault();
    $(this).prop('disabled', false);

    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}/atualizar`,
        method: 'PUT',
        contentType: 'application/json',
        data : {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(res => {
        Swal.fire({
            title: 'Publicação atualizada com sucesso!',
            icon: 'success',
            confirmButtonText: 'OK'
        }).then(() => {
            window.location = "/home";
        });
    }).fail(fail => {
        Swal.fire({
            title: 'Erro ao atualizar publicação',
            text: fail.responseText,
            icon: 'error',
            confirmButtonText: 'OK'
        });
    }).always(() => {
        $(this).prop('disabled', false);
    });
}

function deletarPublicacao(e) {
    e.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Você tem certeza que deseja deletar esta publicação? Esta ação não pode ser desfeita.",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Sim, deletar",
        cancelButtonText: "Cancelar"
    }).then((result) => {
        if (result.isConfirmed) {
            const elementoClicado = $(e.currentTarget);
            const publicacao = elementoClicado.closest('div');
            const publicacaoId = publicacao.data('publicacao-id');
            elementoClicado.prop('disabled', true);
            console.log("Deletando publicação ID: " + publicacaoId);
            $.ajax({
                url: `/publicacoes/${publicacaoId}/deletar`,
                method: 'DELETE'
            }).done(res => {
                publicacao.fadeOut(500, function() {
                    $(this).remove();
                });
            }).fail(fail => {
                Swal.fire({
                    title: 'Erro ao deletar publicação',
                    text: fail.responseText,
                    icon: 'error',
                    confirmButtonText: 'OK'
                });
            }).always(() => {
                elementoClicado.prop('disabled', false);
            });
        }
    });
}