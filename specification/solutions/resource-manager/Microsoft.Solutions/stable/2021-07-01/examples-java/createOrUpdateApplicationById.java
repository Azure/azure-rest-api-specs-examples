import com.azure.resourcemanager.managedapplications.fluent.models.ApplicationInner;

/** Samples for Applications CreateOrUpdateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplicationById.json
     */
    /**
     * Sample code: Creates or updates a managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void createsOrUpdatesAManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .createOrUpdateById(
                "subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication",
                new ApplicationInner()
                    .withKind("ServiceCatalog")
                    .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
                    .withApplicationDefinitionId(
                        "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
                com.azure.core.util.Context.NONE);
    }
}
