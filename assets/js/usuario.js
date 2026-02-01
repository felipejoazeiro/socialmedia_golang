$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('click', editar);

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