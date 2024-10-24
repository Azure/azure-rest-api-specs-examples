using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs.Models;
using Azure.ResourceManager.Avs;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/ScriptExecutions_GetExecutionLogs.json
// this example is just showing the usage of "ScriptExecutions_GetExecutionLogs" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScriptExecutionResource created on azure
// for more information of creating ScriptExecutionResource, please refer to the document of ScriptExecutionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string scriptExecutionName = "addSsoServer";
ResourceIdentifier scriptExecutionResourceId = ScriptExecutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, scriptExecutionName);
ScriptExecutionResource scriptExecution = client.GetScriptExecutionResource(scriptExecutionResourceId);

// invoke the operation
IEnumerable<ScriptOutputStreamType> scriptOutputStreamType = new ScriptOutputStreamType[]
{
ScriptOutputStreamType.Information,new ScriptOutputStreamType("Warnings"),new ScriptOutputStreamType("Errors"),ScriptOutputStreamType.Output
};
ScriptExecutionResource result = await scriptExecution.GetExecutionLogsAsync(scriptOutputStreamType: scriptOutputStreamType);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScriptExecutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
