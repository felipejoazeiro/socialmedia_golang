$("#login").on("submit", fazerLogin);

function fazerLogin(e) {
    e.preventDefault();
    $.ajax({
        url: "/login",
        method: "POST", 
        data: {
            email: $("#email").val(),
            senha: $("#senha").val()
        }
    }).done(function() {
        window.location.href = "/home";
    }).fail(function() {
        Swal.fire({
            title: 'Erro ao fazer login',
            text: 'Verifique suas credenciais.',
            icon: 'error',
            confirmButtonText: 'OK'
        });
    });
}