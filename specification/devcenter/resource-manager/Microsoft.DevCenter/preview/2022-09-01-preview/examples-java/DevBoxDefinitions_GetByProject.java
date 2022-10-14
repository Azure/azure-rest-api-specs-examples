import com.azure.core.util.Context;

/** Samples for DevBoxDefinitions GetByProject. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/DevBoxDefinitions_GetByProject.json
     */
    /**
     * Sample code: DevBoxDefinitions_GetByProject.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devBoxDefinitionsGetByProject(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devBoxDefinitions().getByProjectWithResponse("rg1", "ContosoProject", "WebDevBox", Context.NONE);
    }
}
