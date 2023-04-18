import com.azure.resourcemanager.workloads.fluent.models.SapLandscapeMonitorInner;
import com.azure.resourcemanager.workloads.models.SapLandscapeMonitorMetricThresholds;
import com.azure.resourcemanager.workloads.models.SapLandscapeMonitorPropertiesGrouping;
import com.azure.resourcemanager.workloads.models.SapLandscapeMonitorSidMapping;
import java.util.Arrays;

/** Samples for SapLandscapeMonitor Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_Update.json
     */
    /**
     * Sample code: Update SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void updateSAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapLandscapeMonitors()
            .updateWithResponse(
                "myResourceGroup",
                "mySapMonitor",
                new SapLandscapeMonitorInner()
                    .withGrouping(
                        new SapLandscapeMonitorPropertiesGrouping()
                            .withLandscape(
                                Arrays
                                    .asList(
                                        new SapLandscapeMonitorSidMapping()
                                            .withName("Prod")
                                            .withTopSid(Arrays.asList("SID1", "SID2"))))
                            .withSapApplication(
                                Arrays
                                    .asList(
                                        new SapLandscapeMonitorSidMapping()
                                            .withName("ERP1")
                                            .withTopSid(Arrays.asList("SID1", "SID2")))))
                    .withTopMetricsThresholds(
                        Arrays
                            .asList(
                                new SapLandscapeMonitorMetricThresholds()
                                    .withName("Instance Availability")
                                    .withGreen(90.0F)
                                    .withYellow(75.0F)
                                    .withRed(50.0F))),
                com.azure.core.util.Context.NONE);
    }
}
