/** Samples for Applications CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplication.json
     */
    /**
     * Sample code: Create or update managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void createOrUpdateManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .define("myManagedApplication")
            .withRegion((String) null)
            .withExistingResourceGroup("rg")
            .withKind("ServiceCatalog")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withApplicationDefinitionId(
                "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef")
            .create();
    }
}
