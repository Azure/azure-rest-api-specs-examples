
import com.azure.resourcemanager.datafactory.models.ExposureControlBatchRequest;
import com.azure.resourcemanager.datafactory.models.ExposureControlRequest;
import java.util.Arrays;

/**
 * Samples for ExposureControl QueryFeatureValuesByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ExposureControl_QueryFeatureValuesByFactory.json
     */
    /**
     * Sample code: ExposureControl_QueryFeatureValuesByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        exposureControlQueryFeatureValuesByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.exposureControls().queryFeatureValuesByFactoryWithResponse("exampleResourceGroup", "exampleFactoryName",
            new ExposureControlBatchRequest().withExposureControlRequests(Arrays.asList(
                new ExposureControlRequest().withFeatureName("ADFIntegrationRuntimeSharingRbac")
                    .withFeatureType("Feature"),
                new ExposureControlRequest().withFeatureName("ADFSampleFeature").withFeatureType("Feature"))),
            com.azure.core.util.Context.NONE);
    }
}
