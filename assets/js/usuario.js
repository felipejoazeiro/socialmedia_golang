$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);

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

}