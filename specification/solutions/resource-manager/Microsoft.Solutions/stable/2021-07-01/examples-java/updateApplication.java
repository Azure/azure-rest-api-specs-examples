import com.azure.resourcemanager.managedapplications.fluent.models.ApplicationPatchableInner;

/** Samples for Applications Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplication.json
     */
    /**
     * Sample code: Updates managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void updatesManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .update(
                "rg",
                "myManagedApplication",
                new ApplicationPatchableInner()
                    .withKind("ServiceCatalog")
                    .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
                    .withApplicationDefinitionId(
                        "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
                com.azure.core.util.Context.NONE);
    }
}
