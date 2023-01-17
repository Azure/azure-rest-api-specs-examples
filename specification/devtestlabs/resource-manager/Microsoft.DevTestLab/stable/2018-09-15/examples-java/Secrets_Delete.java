/** Samples for Secrets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Secrets_Delete.json
     */
    /**
     * Sample code: Secrets_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void secretsDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .secrets()
            .deleteWithResponse(
                "resourceGroupName", "{labName}", "{userName}", "{secretName}", com.azure.core.util.Context.NONE);
    }
}
