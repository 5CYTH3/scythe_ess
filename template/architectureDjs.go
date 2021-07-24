package template

/*
func djsProjectContent(rootPath string) {

	file, _ := os.Create(rootPath + "/index.js")

	file.WriteString(
		`import Discord from 'discord.js';
import fs from 'fs';
import { prefix, token } from './config.json';

const client = new Discord.Client();

client.once('ready', () => {
	console.log('Ready!');
});

client.commands = new Discord.Collection();

const commandFiles = fs.readdirSync('./src/commands').filter(file => file.endsWith('.js'));

for (const file of commandFiles) {
	const command = require(`./commands/${file}`);
	client.commands.set(command.name, command);
}`)

}
*/
