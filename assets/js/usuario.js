$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-usuario').on('click', deletarUsuario);

function pararDeSeguir(usuarioId) {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        type: 'POST',
    }).done(() => {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(() => {
        alert('Não foi possível parar de seguir o usuário.');
        $('#parar-de-seguir').prop('disabled', false);
    });
}

function seguir(){
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);
    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        type: 'POST',
    }).done(() => {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(() => {
        alert('Não foi possível seguir o usuário.');
        $('#seguir').prop('disabled', false);
    });
}

function editar(e){
    e.preventDefault();

    const usuarioId = $(this).data('usuario-id');
    const nome = $('#nome').val();
    const email = $('#email').val();
    const senha = $('#senha').val();

    $.ajax({
        url: `/editar-usuario/${usuarioId}`,
        type: 'PUT',
        data: {
            nome: nome,
            email: email,
            senha: senha
        }
    }).done(() => {
        Swal.fire({
            icon: 'success',
            title: 'Usuário editado com sucesso!',
        });
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Erro ao editar usuário.',
        });
    });
}

function atualizarSenha(e){
    e.preventDefault();

    const usuarioId = $(this).data('usuario-id');
    const senhaAtual = $('#senha-atual').val();
    const novaSenha = $('#nova-senha').val();
    const confirmarSenha = $('#confirmar-senha').val();

    if (novaSenha !== confirmarSenha) {
        Swal.fire({
            icon: 'error',
            title: 'A nova senha e a confirmação não coincidem.',
        });
        return;
    }

    $.ajax({    
        url: `/atualizar-senha/${usuarioId}`,
        type: 'PUT',
        data: {
            senhaAtual: senhaAtual,
            novaSenha: novaSenha,
            confirmarSenha: confirmarSenha
        }
    }).done(() => {
        Swal.fire({
            icon: 'success',
            title: 'Senha atualizada com sucesso!',
        }).then (() => {
            window.location = `/perfil`;
        });
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Erro ao atualizar senha.',
        });
    });
}

function deletarUsuario() {
    Swal.fire({
        title: 'Tem certeza que deseja excluir sua conta?',
        text: "Esta ação é irreversível!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Sim, excluir minha conta!',
        cancelButtonText: 'Cancelar'    
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                url: `/deletar-usuario`,
                type: 'DELETE',
            }).done(() => {
                Swal.fire({
                    icon: 'success',
                    title: 'Conta excluída com sucesso!',
                }).then(() => {
                    window.location = `/logout`;
                }
                );
            }
            ).fail(() => {
                Swal.fire({
                    icon: 'error',
                    title: 'Erro ao excluir conta.',
                });
            }
            );
        }
    });
}