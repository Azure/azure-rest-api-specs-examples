
import com.azure.resourcemanager.timeseriesinsights.models.ReferenceDataKeyPropertyType;
import com.azure.resourcemanager.timeseriesinsights.models.ReferenceDataSetKeyProperty;
import java.util.Arrays;

/**
 * Samples for ReferenceDataSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/
     * ReferenceDataSetsCreate.json
     */
    /**
     * Sample code: ReferenceDataSetsCreate.
     * 
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void
        referenceDataSetsCreate(com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.referenceDataSets().define("rds1").withRegion("West US").withExistingEnvironment("rg1", "env1")
            .withKeyProperties(Arrays.asList(
                new ReferenceDataSetKeyProperty().withName("DeviceId1").withType(ReferenceDataKeyPropertyType.STRING),
                new ReferenceDataSetKeyProperty().withName("DeviceFloor")
                    .withType(ReferenceDataKeyPropertyType.DOUBLE)))
            .create();
    }
}
