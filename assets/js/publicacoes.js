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
        alert("Erro ao criar publicação: " + fail.responseText);
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
        alert("Erro ao curtir publicação: " + fail.responseText);
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
        alert("Erro ao descurtir publicação: " + fail.responseText);
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
        window.location = "/home";
    }).fail(fail => {
        alert("Erro ao atualizar publicação: " + fail.responseText);
    }).always(() => {
        $(this).prop('disabled', false);
    });
}

function deletarPublicacao(e) {
    e.preventDefault();
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
        alert("Erro ao deletar publicação: " + fail.responseText);
    }).always(() => {
        elementoClicado.prop('disabled', false);
    });
}