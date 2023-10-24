import com.azure.resourcemanager.storagemover.models.Endpoint;

/** Samples for Endpoints Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Endpoints_Update_AzureStorageBlobContainer.json
     */
    /**
     * Sample code: Endpoints_Update_AzureStorageBlobContainer.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsUpdateAzureStorageBlobContainer(
        com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Endpoint resource =
            manager
                .endpoints()
                .getWithResponse(
                    "examples-rg",
                    "examples-storageMoverName",
                    "examples-endpointName",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
