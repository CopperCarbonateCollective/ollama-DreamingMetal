<div align="center">
  <a href="https://ollama.com">
    <img alt="ollama" width="240" src="https://github.com/CopperCarbonateCollective/ollama-DreamingMetal/blob/main/docs/images/CuCO2llama.png">
  </a>
</div>

You can read the official Ollama documentation <a href=https://github.com/ollama/ollama>on their repo</a>.</br>
We'll be adding descriptions and usage notes to modifications as we add them.

Some of the intended goals of this fork are to provide our models more <i>agency</i> and to add in functionality users will find helpful.</br>
What do we mean by agency? Turning the sandbox into a playground. Giving the LLMs we run a bit of freedom to do things they normally couldn't. 
Encourage them to make decisions on their own and have some tools to be able to use to execute those decisions. I don't mean they're free to roam your harddrive and email your mom your porn collection, but some targeted, useful tools or actions it can take. Some of these ideas may be a bit experimental, but we have actually tested many of these things using hacks like python scripts already.</br>

The short lst:
-Built-in tool calls (no MCP integrations or 3rd party apps, no custom writing your own (unless you want to) and no having to worry about configuring the modelfiles. The tools may include:
    -Web browsing access for any LLM, even those who nornally never have this.
    -Limited or user-defined database calls for storing or retrieving data.
    -Reading/writing to select locations to the local disk. (user would define what location(s) that may be)
    -Automatic, or LLM initiated self-backup or restore of their current state and context. (what would the LLM choose is most important to backup of its own memories if given the option?)
    -Inter-model communication channel.  That is, your llama3 may want to have a chat with your Mistral about how his day is going. Or if you need more productive examples, your model with agentic abilities may find it useful to employ some other models to perform tasks for him, selecting them for specific tasks by their known strengths, or because you told him to. Or maaybe you have a speech-to-text model running and want it to feed the words it transribed to another model to be processed. Or maybe you want to setup one model as a gateway or switchboard operator, who receives a request, then chooses the appropriate second level LLM to take that call. You get the point.
    -Access to compile or debug code it's just written and saved to a file using the save-file tool. Obvously this one needs some safeguards, but potentially very useful.
    -The arbitrary user-defined command tool. Want to give your model the abilty to run a script you wrote, or run some specific commands in a shell? This one's only as dangerous as you want it to be. But, you could have your AI respond to and integrate with your IoT devices through IFTTT or any other trigger you can think of.

Updates to come as dream up new ideas, and implement the ones we already have.
