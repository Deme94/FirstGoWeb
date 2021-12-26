function login(){
    console.log($("#email").val());
    if(!$("#email").val()){
        $("#email").addClass('is-invalid');
        return;
    } else {
        $("#email").removeClass('is-invalid');
    }
    if(!$("#password").val()){
        $("#password").addClass('is-invalid');
        return;
    } else {
        $("#password").removeClass('is-invalid');
    }

    $.post('/loginUser', $('#form').serialize()).done(function (data) {
        if (data=="error"){
            $("#errorLogin").html("Correo electrónico o contraseña incorrectos");
        } else {
            window.location.replace("/admin/usuarios");
        }
    });
}