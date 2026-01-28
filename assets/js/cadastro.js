$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(e) {
    e.preventDefault();
    console.log("Iniciando criação de usuário...");

    var senha = $('#senha').val();
    var confirmarSenha = $('#confirmar-senha').val();

    if (senha !== confirmarSenha) {
        Swal.fire({
            title: 'Erro',
            text: 'As senhas não coincidem!',
            icon: 'error',
            confirmButtonText: 'OK'
        });
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
        Swal.fire({
            title: 'Usuário criado com sucesso!',
            icon: 'success',
            confirmButtonText: 'OK'
        }).then(() => {
           $.ajax({
               url: "/login",
               method: "POST",
                data: {
                    email: $('#email').val(),
                    senha: senha
                }
           }).done(function() {
               window.location.href = "/home";
        }).fail(function() {
               Swal.fire({
                   title: 'Erro ao fazer login automático',
                   text: 'Por favor, faça login manualmente.',
                   icon: 'error',
                   confirmButtonText: 'OK'
               });
           });
        });
    }).fail(fail => {
        Swal.fire({
            title: 'Erro ao criar usuário',
            text: fail.responseText,
            icon: 'error',
            confirmButtonText: 'OK'
        });
    })
}