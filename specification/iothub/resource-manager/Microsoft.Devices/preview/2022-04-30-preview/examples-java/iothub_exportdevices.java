import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.AuthenticationType;
import com.azure.resourcemanager.iothub.models.ExportDevicesRequest;
import com.azure.resourcemanager.iothub.models.ManagedIdentity;

/** Samples for IotHubResource ExportDevices. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_exportdevices.json
     */
    /**
     * Sample code: IotHubResource_ExportDevices.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceExportDevices(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .exportDevicesWithResponse(
                "myResourceGroup",
                "testHub",
                new ExportDevicesRequest()
                    .withExportBlobContainerUri("testBlob")
                    .withExcludeKeys(true)
                    .withAuthenticationType(AuthenticationType.IDENTITY_BASED)
                    .withIdentity(
                        new ManagedIdentity()
                            .withUserAssignedIdentity(
                                "/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")),
                Context.NONE);
    }
}
