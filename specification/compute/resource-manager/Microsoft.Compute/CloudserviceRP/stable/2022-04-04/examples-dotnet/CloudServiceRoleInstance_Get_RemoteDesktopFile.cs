using System;
using System.IO;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudServiceRoleInstance_Get_RemoteDesktopFile.json
// this example is just showing the usage of "CloudServiceRoleInstances_GetRemoteDesktopFile" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this CloudServiceRoleInstanceResource created on azure
// for more information of creating CloudServiceRoleInstanceResource, please refer to the document of CloudServiceRoleInstanceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcloudService";
string cloudServiceName = "aaaa";
string roleInstanceName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier cloudServiceRoleInstanceResourceId = CloudServiceRoleInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServiceName, roleInstanceName);
CloudServiceRoleInstanceResource cloudServiceRoleInstance = client.GetCloudServiceRoleInstanceResource(cloudServiceRoleInstanceResourceId);
            
// invoke the operation
Stream result = await cloudServiceRoleInstance.GetRemoteDesktopFileAsync();
            
Console.WriteLine($"Succeeded: {result}");
