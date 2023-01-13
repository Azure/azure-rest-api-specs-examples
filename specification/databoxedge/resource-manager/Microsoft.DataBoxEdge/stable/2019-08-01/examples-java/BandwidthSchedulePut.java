import com.azure.resourcemanager.databoxedge.models.DayOfWeek;
import java.util.Arrays;

/** Samples for BandwidthSchedules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/BandwidthSchedulePut.json
     */
    /**
     * Sample code: BandwidthSchedulePut.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void bandwidthSchedulePut(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .bandwidthSchedules()
            .define("bandwidth-1")
            .withExistingDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation")
            .withStart("0:0:0")
            .withStop("13:59:0")
            .withRateInMbps(100)
            .withDays(Arrays.asList(DayOfWeek.SUNDAY, DayOfWeek.MONDAY))
            .create();
    }
}
