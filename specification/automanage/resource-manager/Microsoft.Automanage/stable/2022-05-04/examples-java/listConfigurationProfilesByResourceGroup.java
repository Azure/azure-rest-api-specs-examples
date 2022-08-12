import com.azure.core.util.Context;

/** Samples for ConfigurationProfiles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfilesByResourceGroup.json
     */
    /**
     * Sample code: List configuration profiles by resource group.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfilesByResourceGroup(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.configurationProfiles().listByResourceGroup("myResourceGroupName", Context.NONE);
    }
}
