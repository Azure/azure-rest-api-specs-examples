
/**
 * Samples for EnvironmentTypes ListByDevCenter.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/EnvironmentTypes_List.
     * json
     */
    /**
     * Sample code: EnvironmentTypes_ListByDevCenter.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentTypesListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentTypes().listByDevCenter("rg1", "Contoso", null, com.azure.core.util.Context.NONE);
    }
}
