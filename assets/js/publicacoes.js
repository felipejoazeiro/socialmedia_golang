$('#nova-publicacao').on('submit', criarPublicacao);

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