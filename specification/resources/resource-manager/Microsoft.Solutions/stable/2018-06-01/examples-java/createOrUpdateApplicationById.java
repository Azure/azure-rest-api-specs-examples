import com.azure.resourcemanager.managedapplications.fluent.models.ApplicationInner;

/** Samples for Applications CreateOrUpdateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationById.json
     */
    /**
     * Sample code: Create or update application by id.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void createOrUpdateApplicationById(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .createOrUpdateById(
                "myApplicationId",
                new ApplicationInner()
                    .withLocation("East US 2")
                    .withKind("ServiceCatalog")
                    .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG"),
                com.azure.core.util.Context.NONE);
    }
}
