import com.azure.core.util.Context;

/** Samples for ConfigurationProfileAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileAssignmentsBySubscription.json
     */
    /**
     * Sample code: List configuration profile assignments by subscription.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfileAssignmentsBySubscription(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.configurationProfileAssignments().list(Context.NONE);
    }
}
