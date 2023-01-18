/** Samples for Secrets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Secrets_CreateOrUpdate.json
     */
    /**
     * Sample code: Secrets_CreateOrUpdate.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void secretsCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .secrets()
            .define("{secretName}")
            .withRegion((String) null)
            .withExistingUser("resourceGroupName", "{labName}", "{userName}")
            .withValue("{secret}")
            .create();
    }
}
