using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Domains_RegenerateKey.json
// this example is just showing the usage of "Domains_RegenerateKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridDomainResource created on azure
// for more information of creating EventGridDomainResource, please refer to the document of EventGridDomainResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string domainName = "exampledomain2";
ResourceIdentifier eventGridDomainResourceId = EventGridDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName);
EventGridDomainResource eventGridDomain = client.GetEventGridDomainResource(eventGridDomainResourceId);

// invoke the operation
EventGridDomainRegenerateKeyContent content = new EventGridDomainRegenerateKeyContent("key1");
EventGridDomainSharedAccessKeys result = await eventGridDomain.RegenerateKeyAsync(content);

Console.WriteLine($"Succeeded: {result}");
