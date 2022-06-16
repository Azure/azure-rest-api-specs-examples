import com.azure.core.util.Context;

/** Samples for ContainerAppsRevisions GetRevision. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/Revisions_Get.json
     */
    /**
     * Sample code: Get Container App's revision.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerAppSRevision(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsRevisions()
            .getRevisionWithResponse("rg", "testcontainerApp0", "testcontainerApp0-pjxhsye", Context.NONE);
    }
}
