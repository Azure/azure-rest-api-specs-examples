
import com.azure.resourcemanager.datafactory.models.ExposureControlRequest;

/**
 * Samples for ExposureControl GetFeatureValue.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ExposureControl_GetFeatureValue.json
     */
    /**
     * Sample code: ExposureControl_GetFeatureValue.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        exposureControlGetFeatureValue(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.exposureControls().getFeatureValueWithResponse("WestEurope",
            new ExposureControlRequest().withFeatureName("ADFIntegrationRuntimeSharingRbac").withFeatureType("Feature"),
            com.azure.core.util.Context.NONE);
    }
}
