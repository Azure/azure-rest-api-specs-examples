import com.azure.resourcemanager.storagemover.models.Endpoint;

/** Samples for Endpoints Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Endpoints_Update.json
     */
    /**
     * Sample code: Endpoints_Update.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
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
