using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.WeightsAndBiases.Models;
using Azure.ResourceManager.WeightsAndBiases;

// Generated from example definition: 2024-09-18-preview/Instances_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "InstanceResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WeightsAndBiasesInstanceResource created on azure
// for more information of creating WeightsAndBiasesInstanceResource, please refer to the document of WeightsAndBiasesInstanceResource
string subscriptionId = "0BCB047F-CB71-4DFE-B0BD-88519F411C2F";
string resourceGroupName = "rgopenapi";
string instancename = "myinstance";
ResourceIdentifier weightsAndBiasesInstanceResourceId = WeightsAndBiasesInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instancename);
WeightsAndBiasesInstanceResource weightsAndBiasesInstance = client.GetWeightsAndBiasesInstanceResource(weightsAndBiasesInstanceResourceId);

// invoke the operation
await weightsAndBiasesInstance.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
