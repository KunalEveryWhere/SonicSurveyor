#! /bin/bash

# Run the get started turorial
# https://noisemodelling.readthedocs.io/en/latest/Get_Started_Tutorial.html

# Step 4: Upload files to database
# create (or load existing) database and load a shape file into the database
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/ground_type.shp
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/buildings.shp
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/oisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/receivers.shp
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/ROADS2.shp
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/dem.geojson


# Step 5: Run Calculation
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/NoiseModelling/Noise_level_from_traffic.groovy -tableBuilding BUILDINGS -tableRoads ROADS2 -tableReceivers RECEIVERS -tableDEM DEM -tableGroundAbs GROUND_TYPE

# Step 6: Export (& see) the results
./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Export_Table.groovy -exportPath LDAY_GEOM.shp -tableToExport LDAY_GEOM
