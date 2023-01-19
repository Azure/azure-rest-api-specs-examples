import com.azure.resourcemanager.managedapplications.models.Application;

/** Samples for Applications Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplication.json
     */
    /**
     * Sample code: Updates a managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void updatesAManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        Application resource =
            manager
                .applications()
                .getByResourceGroupWithResponse("rg", "myManagedApplication", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withKind("ServiceCatalog")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withApplicationDefinitionId(
                "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef")
            .apply();
    }
}
