/** Samples for ContainerAppsDiagnostics GetRevision. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Revisions_Get.json
     */
    /**
     * Sample code: Get Container App's revision.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerAppSRevision(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsDiagnostics()
            .getRevisionWithResponse(
                "rg", "testcontainerApp0", "testcontainerApp0-pjxhsye", com.azure.core.util.Context.NONE);
    }
}
