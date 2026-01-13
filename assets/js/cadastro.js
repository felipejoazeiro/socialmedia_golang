$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(e) {
    e.preventDefault();
    console.log("Iniciando criação de usuário...");

    var senha = $('#senha').val();
    var confirmarSenha = $('#confirmar-senha').val();

    if (senha !== confirmarSenha) {
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: '/usuarios',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: senha
        })
    }).done(res => {
        alert("Usuário criado com sucesso!");
    }).fail(fail => {
        alert("Erro ao criar usuário: " + fail.responseText);
    })
}