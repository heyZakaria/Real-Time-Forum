export async function Registred() {
    try {
        const response = await fetch("/api/registred", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },

        });

        if (response.status != 200) {
            document.cookie = 'session_id='; 'Max-Age=0'
            


            return false

        } else {
            let x = await response.json();

            return x
        }

       
    } catch (error) {


        console.error('Error fetching data:', error);
        return false
    }

}
