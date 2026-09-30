
/**
 * Samples for CloudValidations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/CloudValidations_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: CloudValidations_ListByResourceGroup_MinimumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void cloudValidationsListByResourceGroupMinimumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.cloudValidations().listByResourceGroup("rgplatformvalidation", null, com.azure.core.util.Context.NONE);
    }
}
