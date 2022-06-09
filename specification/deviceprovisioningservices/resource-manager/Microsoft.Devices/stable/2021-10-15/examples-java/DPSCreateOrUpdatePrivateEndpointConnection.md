```java
import com.azure.resourcemanager.deviceprovisioningservices.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.deviceprovisioningservices.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.deviceprovisioningservices.models.PrivateLinkServiceConnectionStatus;

/** Samples for IotDpsResource CreateOrUpdatePrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSCreateOrUpdatePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_CreateOrUpdate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void privateEndpointConnectionCreateOrUpdate(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .definePrivateEndpointConnection("myPrivateEndpointConnection")
            .withExistingProvisioningService("myResourceGroup", "myFirstProvisioningService")
            .withProperties(
                new PrivateEndpointConnectionProperties()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                            .withDescription("Approved by johndoe@contoso.com")))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceprovisioningservices_1.1.0-beta.1/sdk/deviceprovisioningservices/azure-resourcemanager-deviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.
