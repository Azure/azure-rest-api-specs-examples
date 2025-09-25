using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DisconnectedOperations.Models;
using Azure.ResourceManager.DisconnectedOperations;

// Generated from example definition: 2025-06-01-preview/Artifact_ListDownloadUri_MaximumSet_Gen.json
// this example is just showing the usage of "Artifacts_ListDownloadUri" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DisconnectedOperationsArtifactResource created on azure
// for more information of creating DisconnectedOperationsArtifactResource, please refer to the document of DisconnectedOperationsArtifactResource
string subscriptionId = "301DCB09-82EC-4777-A56C-6FFF26BCC814";
string resourceGroupName = "rgdisconnectedoperations";
string name = "L4z_-S";
string imageName = "B-Ra--W0";
string artifactName = "artifact1";
ResourceIdentifier disconnectedOperationsArtifactResourceId = DisconnectedOperationsArtifactResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, imageName, artifactName);
DisconnectedOperationsArtifactResource disconnectedOperationsArtifact = client.GetDisconnectedOperationsArtifactResource(disconnectedOperationsArtifactResourceId);

// invoke the operation
DisconnectedOperationsArtifactDownloadResult result = await disconnectedOperationsArtifact.GetDownloadUriAsync();

Console.WriteLine($"Succeeded: {result}");
