import{SyntheticEvent, useState} from 'react';

const Message = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [message, setMesaage] = useState('');

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await fetch('http://localhost:8000/api/mesajgonder', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name,
                email,
                message
            })
        });

    }


    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Mesaj Gönder</h1>

            <input className="form-control" placeholder="İsim" required
                   onChange={e => setName(e.target.value)}
            />

            <input type="email" className="form-control" placeholder="Email" required
                   onChange={e => setEmail(e.target.value)}
            />

            <input type="text" className="form-control" placeholder="Mesaj" required
                   onChange={e => setMesaage(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Yolla</button>
        </form>
    );
};

export default Message;
