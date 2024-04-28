var Password = document.getElementById("UserPasswordInput");
var Confirm_Password = document.getElementById("Confirm-PasswordInput");

// Verify if Password is the same as Confirm Password
function PasswordIsEqual(){
    if(Password.value != Confirm_Password.value){
        Confirm_Password.setCustomValidity("Password don't match");
    }else{
        Confirm_Password.setCustomValidity('');
    }
}

Password.onchange = PasswordIsEqual;
Confirm_Password.onkeyup = PasswordIsEqual;

