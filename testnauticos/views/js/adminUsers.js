var register = true;
$(document).ready(function () {

    // Valida los campos de registro (CLIENTE)
    // Validacion Nombre
    $("#name").on('change', function () {
        $("#name").removeClass('is-invalid');
        $("#name").removeClass('is-valid');
        // Regex
        patron = /^([A-Z]{1}[a-zñáéíóú]+[\s]*)+$/;
        match = patron.test($("#name").val());
        if(match){
            $("#name").addClass('is-valid');
        }
        else{
            $("#name").val('');
            $("#name").addClass('is-invalid');
            $("#name-error").html("Nombre incorrecto");
        }
    });

    // Validacion Usuario
    $("#user").on('change', function () {
        $("#user").removeClass('is-invalid');
        $("#user").removeClass('is-valid');
        // Regex
        patron = /^([a-zñáéíóú]+[\s]*)+$/;
        match = patron.test($("#user").val());
        if(match){
            $("#user").addClass('is-valid');
        }
        else{
            $("#user").val('');
            $("#user").addClass('is-invalid');
            $("#user-error").html("Usuario incorrecto");
        }
    });

    // Validacion Email
    $("#email").on('change', function () {
        $("#email").removeClass('is-invalid');
        $("#email").removeClass('is-valid');
        // Regex
        patron = /[\w]+@{1}[\w]+\.[a-z]{2,3}/;
        match = patron.test($("#email").val());
        if(match){
            $("#email").addClass('is-valid');
            validateRegisterServer();
        }
        else{
            $("#email").val('');
            $("#email").addClass('is-invalid');
            $("#email-error").html("Email incorrecto");
        }
    });

    // Validacion Contraseña
    $("#psswd").on('change', function () {
        $("#psswd").removeClass('is-invalid');
        $("#psswd").removeClass('is-valid');
        // Regex
        patron = /[0-9]{8}/;
        match = patron.test($("#psswd").val());
        if(match){
            $("#psswd").addClass('is-valid');
        }
        else{
            $("#psswd").val('');
            $("#psswd").addClass('is-invalid');
            $("#psswd-error").html("Contraseña incorrecta");
        }      
    });

    // Validacion Repetir contraseña
    $("#psswd2").on('change', function () {
        $("#psswd2").removeClass('is-invalid');
        $("#psswd2").removeClass('is-valid');

        isEqual = $("#psswd").val()==$("#psswd2").val();
        if(isEqual){
            $("#psswd2").addClass('is-valid');
        }
        else{
            $("#psswd2").val('');
            $("#psswd2").addClass('is-invalid');
            $("#psswd2-error").html("Contraseña no coincide");
        }         
    });

    // Validacion Telefono
    $("#phone").on('change', function () {
        $("#phone").removeClass('is-invalid');
        $("#phone").removeClass('is-valid');
        // Regex
        patron = /^[\d]{3}[-]*([\d]{2}[-]*){2}[\d]{2}$/;
        match = patron.test($("#phone").val());
        if(match){
            $("#phone").addClass('is-valid');
        }
        else{
            $("#phone").val('');
            $("#phone").addClass('is-invalid');
            $("#phone-error").html("Teléfono incorrecto");
        }        
    });
});

// Valida los campos de registro en la BBDD (SERVIDOR) 
function validateRegisterServer(){
    $.post('/validateRegister', $('#form').serialize()).done(function (data) {
        if(data == "Error"){
            $("#email").removeClass('is-valid');
            $("#email").val('');
            $("#email").addClass('is-invalid');
            $("#email-error").html("Email ya existe");
        }
    });
}

// Registrar o editar usuario
function saveUser(){
    var msg = "";
    var msgError = "";
    if (register){
        msg = "El usuario ha sido registrado con éxito";
        msgError = "El usuario no ha sido registrado";
    } else {
        msg = "El usuario ha sido actualizado con éxito";
        msgError = "El usuario no ha sido actualizado";
    }
    $("#email").prop('disabled', false); //disable 
    $.post('/saveUser', $('#form').serialize()).done(function (data) {
        if(data == ""){
            Swal.fire({
                title: msg,
                icon:'success'
                }
            ).then((result) => {
                location.reload();
            })
        } else{
            Swal.fire({
                title: msgError,
                icon:'error'
                }
            ).then((result) => {
                location.reload();
            })
        }
    });
}

// Elimina un usuario
function deleteUser(id){
    var sendData = { 
        id : id 
    };
    Swal.fire({
        title: '¿Estás seguro de que deseas eliminar este usuario?',
        text: "Esta acción es irreversible",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Sí, borrar usuario',
        cancelButtonText: 'No, cancelar'
      }).then((result) => {
        if (result.isConfirmed) {
            $.post('/deleteUser', sendData).done(function (data) {
                // Eliminar de la BBDD, if eliminado, swal.Fire
                if(data == ""){
                    Swal.fire(
                        'Eliminado',
                        'El usuario ha sido eliminado',
                        'success'
                    )
                    location.reload();
                }
                else{
                    Swal.fire(
                        'Error',
                        'El usuario no ha sido eliminado',
                        'error'
                    )
                }
            });
        }
      })
}

// Activa o desactiva un usuario
function activate(id){
    var sendData = { 
        id : id 
    };
    $.post('/activateUser', sendData).done(function (data) {
        if (data==""){
            $("#activate"+id).removeClass('badge-danger');
            $("#activate"+id).addClass('badge-success');
            $("#activate"+id).html("Activado");
        }else{
            $("#activate"+id).removeClass('badge-success');
            $("#activate"+id).addClass('badge-danger');
            $("#activate"+id).html("Desactivado");
        }
    });
}

function openRegisterUser(){
    register = true;
    clearForm();
    $('#form').val('0');
    $('.modal-title').html('Registrar usuario');
    $("#email").prop('disabled', false); //enable
}
// Abre el modal de edicion y carga los campos de un usuario (no lo actualiza en la BBDD)
function openEditUser(id){
    register = false;
    $(".modal-title").html("Editar usuario");
    var sendData = { 
        id : id 
    };
    $("#email").prop('disabled', true); //disable 
    $.get('/getUser', sendData).done(function (data) {
        var dataJson = JSON.parse(data);
        $("#name").val(dataJson.Name);
        $("#user").val(dataJson.Username);
        $("#email").val(dataJson.Email);
        $("#psswd").val(dataJson.Password);
        $("#psswd2").val(dataJson.Password);
        $("#phone").val(dataJson.Phone);
        $(".js-example-basic-multiple").val(dataJson.Courses);
        $('.js-example-basic-multiple').trigger('change'); 
    });
}

function clearForm(){
    $('#form')[0].reset();
    $("#name").removeClass('is-invalid');
    $("#name").removeClass('is-valid');
    $("#user").removeClass('is-invalid');
    $("#user").removeClass('is-valid');
    $("#email").removeClass('is-invalid');
    $("#email").removeClass('is-valid');
    $("#psswd").removeClass('is-invalid');
    $("#psswd").removeClass('is-valid');
    $("#psswd2").removeClass('is-invalid');
    $("#psswd2").removeClass('is-valid');
    $("#phone").removeClass('is-invalid');
    $("#phone").removeClass('is-valid');
    $('.js-example-basic-multiple').val(null).trigger('change');

}

