import com.azure.resourcemanager.operationsmanagement.fluent.models.ManagementAssociationInner;
import com.azure.resourcemanager.operationsmanagement.models.ManagementAssociationProperties;

/** Samples for ManagementAssociations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementAssociationCreate.json
     */
    /**
     * Sample code: SolutionCreate.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void solutionCreate(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager
            .managementAssociations()
            .createOrUpdateWithResponse(
                "rg1",
                "providerName",
                "resourceType",
                "resourceName",
                "managementAssociation1",
                new ManagementAssociationInner()
                    .withLocation("East US")
                    .withProperties(
                        new ManagementAssociationProperties()
                            .withApplicationId(
                                "/subscriptions/sub1/resourcegroups/rg1/providers/Microsoft.Appliance/Appliances/appliance1")),
                com.azure.core.util.Context.NONE);
    }
}
