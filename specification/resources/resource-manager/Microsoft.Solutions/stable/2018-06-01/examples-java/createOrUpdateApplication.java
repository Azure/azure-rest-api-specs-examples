/** Samples for Applications CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplication.json
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
            .withRegion("East US 2")
            .withExistingResourceGroup("rg")
            .withKind("ServiceCatalog")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .create();
    }
}
