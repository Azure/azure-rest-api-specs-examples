using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Chaos.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Chaos;

// Generated from example definition: 2025-01-01/Experiments_Cancel.json
// this example is just showing the usage of "Experiments_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ChaosExperimentResource created on azure
// for more information of creating ChaosExperimentResource, please refer to the document of ChaosExperimentResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string experimentName = "exampleExperiment";
ResourceIdentifier chaosExperimentResourceId = ChaosExperimentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, experimentName);
ChaosExperimentResource chaosExperiment = client.GetChaosExperimentResource(chaosExperimentResourceId);

// invoke the operation
await chaosExperiment.CancelAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
