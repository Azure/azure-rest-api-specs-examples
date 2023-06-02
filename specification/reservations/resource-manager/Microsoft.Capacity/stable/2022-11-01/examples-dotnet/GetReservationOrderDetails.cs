using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationOrderDetails.json
// this example is just showing the usage of "ReservationOrder_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReservationOrderResource created on azure
// for more information of creating ReservationOrderResource, please refer to the document of ReservationOrderResource
Guid reservationOrderId = Guid.Parse("a075419f-44cc-497f-b68a-14ee811d48b9");
ResourceIdentifier reservationOrderResourceId = ReservationOrderResource.CreateResourceIdentifier(reservationOrderId);
ReservationOrderResource reservationOrder = client.GetReservationOrderResource(reservationOrderResourceId);

// invoke the operation
ReservationOrderResource result = await reservationOrder.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReservationOrderData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
