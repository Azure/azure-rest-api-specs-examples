import com.azure.resourcemanager.managedapplications.fluent.models.ApplicationInner;

/** Samples for Applications UpdateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplicationById.json
     */
    /**
     * Sample code: Update application by id.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void updateApplicationById(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .updateByIdWithResponse(
                "myApplicationId",
                new ApplicationInner()
                    .withKind("ServiceCatalog")
                    .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
                    .withApplicationDefinitionId(
                        "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
                com.azure.core.util.Context.NONE);
    }
}
