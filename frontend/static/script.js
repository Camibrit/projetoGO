document.addEventListener('DOMContentLoaded', () => {
    const sobre = document.querySelector('#sobre');
    const comoFunciona = document.querySelector('#como-funciona');
    const btnDoar = document.getElementById('btnDoar');
    const formModal = document.getElementById('formModal');
    const form = document.getElementById('formDoacao');
    
    // Tooltips elegantes
    sobre.addEventListener('mouseover', (e) => {
        const tooltip = document.createElement('div');
        tooltip.className = 'tooltip';
        tooltip.textContent = 'Somos uma instituição comprometida em fazer a diferença. Há mais de 5 anos, trabalhamos incansavelmente para levar alimentos e esperança a pessoas em situação de vulnerabilidade social. Nossa missão é transformar vidas através da solidariedade.';
        document.body.appendChild(tooltip);
        
        const rect = e.target.getBoundingClientRect();
        tooltip.style.left = `${rect.left}px`;
        tooltip.style.top = `${rect.bottom + 5}px`;
        tooltip.style.display = 'block';
        
        sobre.addEventListener('mouseout', () => tooltip.remove());
    });

    comoFunciona.addEventListener('mouseover', (e) => {
        const tooltip = document.createElement('div');
        tooltip.className = 'tooltip';
        tooltip.textContent = 'É simples! Clique em Fazer Doação, preencha o formulário com os itens que deseja doar. Nossa equipe fará a coleta e distribuição aos mais necessitados. Você também pode acompanhar o status da sua doação através do nosso sistema.';
        document.body.appendChild(tooltip);
        
        const rect = e.target.getBoundingClientRect();
        tooltip.style.left = `${rect.left}px`;
        tooltip.style.top = `${rect.bottom + 5}px`;
        tooltip.style.display = 'block';
        
        comoFunciona.addEventListener('mouseout', () => tooltip.remove());
    });
    contato.addEventListener('mouseover', (e) => {
        const tooltip = document.createElement('div');
        tooltip.className = 'tooltip';
        tooltip.innerHTML = `
            📞 (11) 99999-9999<br>
            📧 contato@doarmais.org<br>
            📍 Av. da Solidariedade, 123
        `;
        document.body.appendChild(tooltip);
        
        const rect = e.target.getBoundingClientRect();
        tooltip.style.left = `${rect.left - 200}px`; // Movido 200px para a esquerda
        tooltip.style.top = `${rect.bottom + 5}px`;
        tooltip.style.display = 'block';
        
        contato.addEventListener('mouseout', () => tooltip.remove());
    });
    
    btnDoar.addEventListener('click', () => {
        formModal.style.display = 'block';
    });

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        formModal.style.display = 'none';
        alert('Doação registrada com sucesso! Agradecemos sua contribuição.');
        form.reset();
    });

    window.addEventListener('click', (e) => {
        if (e.target === formModal) {
            formModal.style.display = 'none';
        }
    });
});