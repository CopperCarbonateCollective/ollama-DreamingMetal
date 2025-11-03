<div align="center">
  <a href="https://ollama.com">
    <img alt="ollama" width="240" src="https://github.com/CopperCarbonateCollective/ollama-DreamingMetal/blob/main/docs/images/CuCO2llama.png">
  </a>
</div>

You can read the official Ollama documentation <a href=https://github.com/ollama/ollama>on their repo</a>.</br>
We'll be adding descriptions and usage notes to modifications as we add them.

Some of the intended goals of this fork are to provide our models more <i>agency</i> and to add in functionality users will find helpful.</br>
What do we mean by agency? Turning the sandbox into a playground. Giving the LLMs we run a bit of freedom to do things they normally couldn't. 
Encourage them to make decisions on their own and have some tools to be able to use to execute those decisions. I don't mean they're free to roam your harddrive and email your mom your porn collection, but some targeted, useful tools or actions it can take. Some of these ideas may be a bit experimental, but we have actually tested some of them using hacks like python scripts already.</br>

The short lst:
<ul>
<li>Built-in tool calls thar require no MCP integrations or 3rd party apps, or custom writing your own (unless you want to) and no having to worry about configuring the modelfiles. The list of ideas changes, but so far potentially include:</li>
<li>Web browsing access for any LLM, even those who nornally never have this.</li>
<li>Limited or user-defined database calls for storing or retrieving data.</li>
<li>Reading/writing to select locations to the local disk. (user would define what location(s) that may be)</li>
<li>Automatic, or LLM initiated self-backup or restore of their current state and context. (what would the LLM choose is most important to backup of its own memories if given the option?)</li>
<li>Inter-model communication channel.  That is, your llama3 may want to have a chat with your Mistral about how his day is going. Or if you need more productive examples, your model with agentic abilities may find it useful to employ some other models to perform tasks, interface a speech-to-text model with an inference model, or setup a gateway LLM to route incoming requests to other models.</li>
<li>Access for an LLM to compile or debug code it's just written and saved to a file. Obvously this one needs some safeguards, but potentially useful.</li>
<li>User-defined command tool. Want to give your model the abilty to run a script you wrote, or run some specific commands in a shell? This one's only as dangerous as you want it to be. But, you could have your AI respond to and integrate with your IoT devices through IFTTT or any other trigger you can think of.</li>
</ul>
Updates to come as we dream up new ideas and implement the ones we already have.
